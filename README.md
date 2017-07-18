# Overview

go-word-count reads list of urls from stdin, loads page content and counts number of 'Go' string for each loaded page. 

# Install
Put go-word-count source directory inside it. Example:

    $ git clone https://github.com/agorobets/go-word-count src/go-word-count #
    $ cd src/go-word-count

To compile and run app execute:

    $ echo -e "<list of urls>" | make run

