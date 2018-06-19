# TGIK Documentation 

This will mostly be internal, but sharing broadly as a resource. 


# Scheduling a new episode

### Creating the graphic 

In the internal team Marketing drive find the TGIK directory.
Duplicate the most recent .ai file and rename to the next number. 

*Example:*

```
Duplicate tgik042.ai
Rename to tgik043.ai
```

Find an image to use that is fun, and subtly hints at the topic for the episode. 
For example [episode 38](https://github.com/heptio/tgik/tree/master/episodes/038) was on Kata containers, and the [supporting image was of containers in a shipyard](https://github.com/heptio/tgik/blob/master/episodes/038/038.png).
Use the eyedropper tool while holding shift in Adobe Illustrator to change the accent color to a contrasting color from the image.
Save as a .png

*Example:*

```
tgik043.ai ---> 043.png
```

### Scheduling on YouTube

Go to the [YouTube creator studio](https://www.youtube.com/my_live_events)

Create a [new Live event](https://www.youtube.com/my_live_events?action_create_live_event=1)

```
title: TGI Kubernetes 001: A great name for my episode
description: Come hang out with Kris Nova as she does a bit of hands on hacking of Kubernetes and related topics. Some of this will be Kris talking about the things she knows. Some of this will be Kris exploring something new with the audience. Come join the fun, ask questions, comment, and participate in the live chat!
```

Create the event (Click save)

Update the event with the image you created

Find the YouTube live stream URL of the event

### Creating the links

Log into Rebrandly and create a new link pointing to the YouTube live stream.

*Example:*

```
j.hept.io/tgik001 ---> youtube.com/watch?v=ixs1-UnWiGU
```

Update the pointer pointer for the broad "Watch Live Now!" link. (This allows us to hard code the link in various without having to push a code change).

*Example:*

```
j.hept.io/tgik-live ---> j.hept.io/tgik001 ---> youtube.com/watch?v=ixs1-UnWiGU
```

### Tweet!

Tweet out announcing the episode with the hashtag `#TGIK8s`
Include the rebrandly URL generated above
Include the image created above

*Example:*

<p align="center"><img src="https://github.com/heptio/tgik/blob/master/example-tweet.png" width="360"></p>

### Schedule reminder tweets from the Heptio twitter account

Log into Hubspot

From the top menu bar:

Social -> Publishing

[Create social post]

Create 2 tweets using similar verbage as the tweet created above but in 3rd person.

(Wed or Thurs) reminder tweet
Fri reminder tweet

Schedule them for around 9-11am Pacific time


