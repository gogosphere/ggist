
[godoc](https://godoc.org/github.com/gogosphere/ggistlib)
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
***ggistPrompt="[Y|N]"*** # 'N' to skip prompting for "ready to paste"

I'll be adding a few others over time like don't prompt before pulling fro paste buffer, alternative gist server, etc. I personally use my .bashrc file to manage all my environment so in it I have the following:

```
export ggistToken="$the token you created"
export ggistUser="$your github account name"
export ggistPrompt='N'
```

### Installing

### The go way
```
git clone https://github.com/gogosphere/ggist.git
cd ggist
(make sure your GOPATH is working)
go get -v
... wait for packages to get done ...
go run ./main.go 
```

### Rest of the guide

Not much to it, go to releases ([https://github.com/gogosphere/ggist/releases]()) and download the latest binary that matches your operating system (Only Darwin at the moment).  No switches are required to be used for this to run.  If you exec ggist without any switches it will show a list of all your gists.

```
type to pointer
Pull: https://gist.github.com/2c34410ac37d9f1b97fafa0449cb538d.git

generate hash from string, golang
Pull: https://gist.github.com/f804f3a832dd45d647006f2bc907f84b.git

Go code for running a data dog query
Pull: https://gist.github.com/83536cdb7e7e8b45f3b315c45c500d03.git

```

If you exev ggist with the -s switch you can send any combination of the other switches, other wise ggist will use default descrptions and a hash-generated filename which you can change at you leasure in the gist UI at github.

```
ggists whancock$ ggist -h
Usage: [parameter] command
-h, --help                   Show this help
-s, --send                   send gist
-p, --public                 send to PUBLIC gist, default is PRIVATE
-d, --description=VAL        Twitterish length description of your gist
-f, --file=VAL               Filename, this will be the filename if cloned
-v, --version                Print version and exit
```

### Here's the greatest part!
IT PULLS FROM YOUR SYSTEM CLIPBOARD!!

```
ggists whancock$  ggist -s
Is your clipboard ready? Press [ENTER] when ready to send..
URL: https://gist.github.com/1ffce1097e4f4d5db326a55dea042577.git
```

Again, other switches for (-d)esription (-f)ilename and (-p)rivate are not required but you can use them in any combination if you wish to save time later.  The app defaults to a private gist but sending ***ggist -s -p*** will send your gist to the public side for sharing.  When the command completes it will return the clone URL which is perfectly suitable for sharing (as long as you set the gist to public).



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

