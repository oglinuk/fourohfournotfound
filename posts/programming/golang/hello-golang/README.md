# Hello Golang!

In one of the classes during my bachelor of IT, I started making my first program in [Golang](https://en.wikipedia.org/wiki/Go_(programming_language)). That first program is what is now know as [goccer](https://github.com/OGLinuk/goccer). The idea behind the program was inspired by my fascination of [web crawling](https://en.wikipedia.org/wiki/Web_crawler). This blog is not going to cover every part of goccer, but will cover the key components. It will also cover the issues I ran into and the solution(s) that I used to overcome the problem(s).

The first of the key components is the configuration of the crawler. This is done with a [JSON](https://en.wikipedia.org/wiki/JSON#Example) file. In my case the structure of the `config.json` file is as follows.

```Go
type Config struct {
	MaxWorkers int
	Seeds      []string
}
```

It's pretty simple, there is an int value called `MaxWorkers`, which is essentially the number of cores that the host machine has and the number of workers it can run. `Seeds` is a string array containing the urls to initially crawl.

The next key components are the producer/consumer. One of the first iterations of [goccer](https://github.com/OGLinuk/goccer) crawled urls sequentially (in a for loop). If you've never heard of [time complexity](https://en.wikipedia.org/wiki/Time_complexity), I would suggest spending some time to study and grasp a basic understanding of the concept. In my case I ran into O(n), and usually by the second or third time running [goccer](https://github.com/OGLinuk/goccer) the number of seeds grew to > 100000. Enter the [producer/consumer](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem) pattern. My friend [Sam](https://github.com/pigeonhands) sent me a blog post on [handling 1 million requests per minute with Go](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/), which is what I went off of originally. The following is the current producer/consumer implementation.

```Go
// Producer

func InitProducer(cfg Config) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= cfg.MaxWorkers; i++ {
		go consume(jobs, wg)
	}

	for i, seed := range cfg.Seeds {
		wg.Add(1)
		go func(i int, seed string) {
			log.Printf("Fetching[%d]: %s", i, seed)
			jobs <- Job{URL: seed}
		}(i, seed)
	}
	wg.Wait()
	close(jobs)
}
```

```Go
// Consumer

type Job struct {
	URL string
}

func consume(jobs <-chan Job, wg *sync.WaitGroup) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			q := hive.NewQueen(job.URL, ArchiveFile)
			q.SpawnDrone()
			q.Aggregate()
			wg.Done()
		}
	}
}
```

The concept (at least how I understand it) is pretty straightforward. The producer creates jobs and sends them to a channel, and the consumer takes the jobs from the channel and creates a new crawler that crawls, then aggregates the results. I used to play zerg in [Starcraft 2](https://en.wikipedia.org/wiki/StarCraft_II%3A_Wings_of_Liberty), hence the naming convention queen and drone.

This brings me to the last of the key components; the queen. Originally I was writing all extracted urls to a single file, which I would scan for duplicates. This spawned two problems that I had to fix. The first was `ioutil.ReadFile`. In the first iterations of [goccer](https://github.com/OGLinuk/goccer) I used `ioutil.ReadFile`, which was fine until the file grew to > 1GB (which happened quite fast). This is because `ioutil.ReadFile` reads the contents of the file into memory. This was solved by using `os.Open` in combination with `bufio.Scanner`. The second problem I faced was using a single file. Trying to check for duplicates when the file grew to > 1GB causes another O(n) issue. To solve this issue [Sam](https://github.com/pigeonhands) suggested creating a file for each domain name and appending any paths to the file. The following is the implementation of that concept.

```Go
type URLFile struct {
	file *os.File
	urls map[string]struct{}
}

type URLWriter struct {
	path     string
	urlFiles map[string]*URLFile
}
```

The `URLFile` file is the file created for the domain name, and the map is the paths within the domain. The reasoning for the value of the map being `struct{}` is because an [empty struct has a size of 0](https://dave.cheney.net/2014/03/25/the-empty-struct), and the reason for using a map is the lookup time is O(1). 

The `URLWriter` is a directory containing the files, and again I use a map (to check for duplicates) because the lookup time is O(1).

I didn't cover a few of the components like aggregation, parsing, or archiving, but feel free to browse and read the code if you're interested.
