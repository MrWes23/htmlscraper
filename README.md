# About

This is a simple implementation of a Go script that can scrape any websites response body and create files of this data. The script was written with 3 websites hard coded, but in reality can be a vast amount more. If more websites are added it's possible you will run into runtime speed issues, but with Go's goroutines it helps facilitate this problem.


The first step begins by first initializing the URLs to be used inside of the urls slice. This variable will be used in the for loop as an iterator. Next I create the channel that will be used by the htmlScrape function to send out progress/status notifications to the terminal. These notifications include the request status code and/or any errors that may occur among other statuses.

The for loop is then constructed which will in turn be the location from which the goroutines spin up. Moving on to the function, it begins by sending a get request to the specified url in which either a 200 status code returns or an error. When the 200 status code is received the function moves on to the else statement in which the ioutil package is used to read the body of the request and write to file as well. If the code is run as written, it will create 3 files from each website containing all of the response body HTML.

