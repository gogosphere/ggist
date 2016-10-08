
[godoc](https://godoc.org/github.com/gogosphere/ggist)

# GGist
I was constantly on the lookout for places to save code snippets in a tool that was not meant to be a catch all.  A place that could do some light syntax highlighting.  I tried Evernote (love it for general catch all notes), I tried OneNote (I don't love it), and various other things like Codebox, Geeknote, etc.  I know there had been gist clis written and they are handy enough but I wanted something that was fast, required little to no switches and arg and gets out of its own way.  

## Getting Started

To use it, go to releases and download the binary.  To hack on it, clone, hack, pull.

### Prerequisities

You need to have a github account, and generate a Personal Access Token [(https://github.com/settings/profile)]() and the only permissions required is 'gists'.

![gist permissions image](https://cloud.githubusercontent.com/assets/20846757/17717626/b602d3f8-63d4-11e6-9091-925be78d90aa.png)

I don't like config files for command line tools, it seems somehow cluncky and less portable that environment variables.  So, set an environment variable to:

***ggistToken="$the token you created"***
***ggistUser="$your github account name"***


```
export ggistToken="$the token you created"
export ggistUser="$your github account name"
```

### Installing

### The go way
```
git clone https://github.com/gogosphere/ggist.git
cd ggist
(make sure your GOPATH is working)
go get -d
... wait for packages to get done ...
go run ./main.go 
```

### Rest of the guide

Not much to it, go to releases ([https://github.com/gogosphere/ggist/releases]()) and download the latest binary that matches your operating system (Only Darwin at the moment).  No switches are required to be used for this to run.  If you exec ggist without any switches it will show a list of all your gists.

```
prompt$ ggist
generate hash from string, golang
Pull: https://gist.github.com/f804f3a832dd45d647006f2bc907f84b.git

Go code for running a data dog query
Pull: https://gist.github.com/83536cdb7e7e8b45f3b315c45c500d03.git

```

To send a gist, select the data you wish to transmit and copy it to your machines clipboard.  [CMD-C]  If you either [-s]et or [-m]atch you must pass the [-d]escription followed by either the description or summary of your gist or an explict match word to search for.

```
ggist -h
Usage: [parameter] command
-h, --help                   Show this help
-s, --set                    send gist (exclusive)
-g, --get                    get gist and display (exclusive) [the default option if no swtiches are specified]
-m, --match                  get gist and display
-v, --debug                  debuging output
-d, --description=VAL        Twitterish length description of your gist
```

### Here's the greatest part!

```
ggists whancock$  ggist -s -d "this is a summary of the gist which I can search on later"
URL: https://gist.github.com/1ffce1097e4f4d5db326a55dea042577.git
```

When the command completes it will return the clone URL which is perfectly suitable for sharing (as long as you set the gist to public).


## Running the tests

Haven't gotten this far.

### Want to beta?
Pull the binary and run it or clone the repo and 

```
go run ggist.go 
```


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Author

* **Bill Hancock** - *Initial work* - [Gogosphere](https://github.com/gogosphere)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.. or whatever, just give me props if you like it.

## Acknowledgments

* Frustration
* Not getting what I want out of things
* iMAC 27"

