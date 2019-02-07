# TGIK Documentation

This will mostly be internal, but sharing broadly as a resource.

## Checklists

Start by copying the following into the issue for the episode.  Check things off as you do them.

```markdown
The process for producing an episode is [here](https://github.com/heptio/tgik/blob/master/documentation/README.md).  All the detail for the steps below is there.

Monday or Tuesday:
* [ ] Pick a topic and create/use an issue in the repo
* [ ] Create the graphic.
* [ ] Schedule on Youtube.
* [ ] Add scheduled episode to TGIK playlist on YT
* [ ] Create short link in rebrandly
* [ ] Tweet about it from both personal and company accounts
* [ ] Schedule future tweets
* [ ] Post tweets on internal #tgik slack channel so that others know.

Before episode:
* [ ] Create HackMD page for live notes.  Start with README template. (TODO: @castrojo -- put details here?)
* [ ] Get everything tested well before episode so you can debug.
* [ ] Use internal slack channel to ensure that AV setup is good

After show (ideally right away but can wait until Monday):
* Create/submit PR for this repo
  * [ ] Episode directory and README from hackmd
  * [ ] Any other files from the episode that might be useful
  * [ ] Update playlist.md with episode
  * Feel free to self merge
* [ ] Edit YT description to point to episode directory here
```

## Scheduling a new episode

### Creating the graphic

Find an image to use that is fun, and subtly hints at the topic for the episode.
For example [episode 38](https://github.com/heptio/tgik/tree/master/episodes/038) was on Kata containers, and the [supporting image was of containers in a shipyard](https://github.com/heptio/tgik/blob/master/episodes/038/038.png).

General process:
* Copy a previous episode into a new file and renumber it
* Pick out an image (often we use unsplash.com). Make sure we have rights to use the image.
* Open the file in Adobe Illustrator
* Delete the old image and use the "File | Place" to insert the new image. Move that to the back with "Object | Arrange | Send to Back"
* Select the semi-transparent background on the text and the logo box and use the eye dropper tool to pick a light color out of the image.
* Export to PNG in the same folder using "File | Export | Export for screens"
  * Before saving make sure you change the name of the artboard (which determines the filename) by double clicking it on that dialog.
  * Export as PNG at 1x

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

For now, use Tweetdeck to schedule tweets.
Create 2 tweets using similar verbiage as the tweet created above but in 3rd person.
Do these for Wed/Thurs and Friday.
Schedule them for around 9-11am Pacific.
Use the j.hept.io link and attach the image.



