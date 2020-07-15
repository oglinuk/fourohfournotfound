---
title: "INotebook"
date: 2018-11-20
tags: ["Python,", "Cryptography,", "Vigenere Cipher"]
draft: false
---

If you have seen the Prestige, then this program will look familiar. INotebook is a personal notebook that is encrypted with the vigenere cipher. A vigenere cipher is similar to a caesar cipher (where characters are shifted by a certain offset), but the vigenere cipher uses a key to determine the offset rather than a number. The key is repeated to the length of the message, then each character of the message is offset by the character of the key. For example if I had a key of hello and a message of world; w would be offset by 8 (since h is 8th in position), e would be offset by 15, and so on. In that example the message is the exact length of the message, so the key wouldnlt be repeated but for messages larger than the hey length, the key is repeated. An example of this would look like

```
key = 'HELLO'
message = ' MESSAGE'

offsets = (['M', 'H'], ['E', 'E',], ['S', 'L'], ['S', 'L'], ['A', 'O'], ['G', 'H'], ['E', E'])
```

Now to write some Python ...

The first thing is to make a class which will be our object for the notebook. This looks like the following ...

```Python
class IllusionistsNotebook(object):
	def __init__(self, key):
		self.key = self.mapKey(key)
        	self.matrix = [ascii_uppercase[(i + 1) % 26] for i in range(len(ascii_uppercase))]
		self.entriesFile = 'entries.txt'
		self.entries = []
	
	def mapKey(self, k):
		key = ''
		for i, c in enumerate(k):
			if c.isalpha():
				key += c.upper()
			if c.isdigit():
				key += ascii_uppercase[int(c) % len(ascii_uppercase)].upper()
		return key
```

The class takes a key which is passed to the mapKey function. The purpose of the mapKey function is to be able to use both alphabet characters and numerical values (so that I can use hash's produced by my personal SBH program). `Self.matrix` is an array of all uppercase alphabetical characters to map the message characters to. The `self.entriesFile` value is a hardcoded file that is used to store entries of inotebook, and `self.entries` is for the entries made within the current session. The next function is the vigenere encryption function, which takes a string as a parameter. The function creates an empty string for which the resulting values are appended to. `repK` is a variable that maps the key to the entire message (repeating the key characters if the message is longer than the key). The next thing that is done is the actual mapping of the individual characters in the passed string to the key characters. ~ Remember the example above? The code itself is pretty simple.

```Python
for x, i in enumerate(s.upper()):
	if i not in ascii_uppercase:
		cipherText += i
	else:
		cipherText += ascii_uppercase[(ascii_uppercase.find(i) + self.matrix.index(repK[x])) % 26]
	return cipherText
```

For each character that is enumerated over, the program checks to see if the current character is within the ascii_uppercase variable, if not the program will append the (i)ndex value of the ascii_uppercase variable. If the character is *not* within the ascii_uppercase variable then the program appends the value of i. I am not entirely sure if this is bad practice since there is essentially nothing changed to a non-alphabetical character. I will leave this for now since it is not a massive priority, but will be something that I will adjust in the future. I already have an idea for a solution, which does the same thing as the else statement above, but for numerical values and for symbols. 

