# Episode 150 : "Production Kubernetes" Book Discussion

- Hosted by @jbeda, @krisnova/@kris-nova
- Recording date: 2021-03-25 

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=blanHrC0SFQ" target="_blank"><img src="https://i.ytimg.com/vi/blanHrC0SFQ/maxresdefault.jpg" border="10" /></a>

This week we are trying something new on TGIK! Joe is inviting the authors of "Production Kubernetes" on to talk about the book.  Come, hang out and ask questions.  Let's see how many of Josh, Rich, Alex and John we can get to join us.

* You can get a FREE (in exchange for email) copy here: https://tanzu.vmware.com/content/ebooks/production-kubernetes
* You can also buy a paper copy book: https://www.amazon.com/Production-Kubernetes-Successful-Application-Platforms/dp/1492092304

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:03:30 - Week in Review
- 00:19:35 - Start if panel/interview

## Week in Review

 - All the container ship memes
     - [Thread on pilot's perspective](https://twitter.com/Nature_Grrrl/status/1375168720495075338?s=20)
     - isitstillstuck.com is a great aggregator
 - Blog post: [Using kubernetes custom resources to manage our ephemeral environments](https://medium.com/beamdental/using-kubernetes-custom-resources-to-manage-our-ephemeral-environments-f298610893e1)
 - [crun](https://github.com/containers/crun) nóva found this week? A fast and low-memory footprint OCI Container Runtime fully written in C.
 - [linux 5.12 RC](http://lkml.iu.edu/hypermail/linux/kernel/2103.0/06524.html) (nóva: Feel free to remove this)
     - [Eyes on 5.12](https://arstechnica.com/gadgets/2021/03/psa-linux-folks-stay-away-from-the-5-12-rc1-kernel/) `linux-5.12 RC1` renamed to `dontuse` (SWAP is broken)
     - [Linux on Nintendo64](https://www.phoronix.com/scan.php?page=news_item&px=Nintendo-64-Controller-Driver)
     - [Kubernetes' swap issue](https://github.com/kubernetes/kubernetes/issues/53533)
       - [Design document from Elana Hashman](https://docs.google.com/document/d/1CZtRtC8W8FwW_VWQLKP9DW_2H-hcLBcH9KBbukie67M/edit#heading=h.alvjn0omorwh) for swap support in K8s 

## Show Notes
- [autopilot whitepaper](https://dl.acm.org/doi/pdf/10.1145/3342195.3387524)
- _Production Kubernetes_ Book Club DM or Slack @csantanapr
## Questions for the Authors

 - [X] Tell us a bit about yourselves.  What led you to this?
 - [X] It takes a lot of personal commitment to write a book. When did you decide to become a technical author and would you go back and do it agian?
 - [X] What is your favorite part of the book?
     - [X] Least favorite?
 - [X] Carlos Santana: Talk about the cover. For a second in twitter I thought it was the Kubernetes Up and Running a Whale vs Dolphin
 - [X] What are some of the biggest pitfalls as folks look to take k8s to production?
 - Peter Lindblom: what do you guys think. big shared cluster or single cluster per tenant?
 - [X] What are some topics that you had to leave out of the book for whatever reason?
 - [ ] Are there concerns of the content becoming stale / outdated?
 - [X] Would you do it again?
 - [X] Thoughts to aspiring authors?
 - [ ] Is there anything unexpected that came out of the writing process?
 - [X] Anything you would do differently? 
 - [K] Anything you wish somebody would have told you?
 - [K] Four authors is a lot! How did you organize splitting up the effort?
 - [X] Duffie: When do you think the benefit of using aggregated apis becomes valuable in comparison to having all the resources defined as extensions to the existin apiserver? There is some real value in using a single "controller manager" for multiple api groups. How does that factor in?
 - [X] Sevi Karaköse: How do you get to write a book with a team of authors? During a world wide pandemic?
 - [ ] Carlos Santana: How long it took to write? What was the timeline?
 - [ ] Smalls: Did Josh coordinate the work in Miro? ;) (because he does EVERYTHING in Miro!)
 - [ ] Alex: How does this book compare or differ from other books like it. For example "Kubernetes and Docker -- An Enterprise Guide" by Surovich and Boorstein.
 - [ ] how to create and test network policies in a microservices environment? We cannot test network policy just with 1 microservice keeping loose coupling concept in mind...
 - [X] Smalls: Given the size of the book and the reputation of K8s being hard (Zulfikar asks similar in chat) what's your opinion of GCP Autopilot or similar? (Not disparaging Autopilot here, just curious)
 - [X] Looking back, what is something you wish you/we/us would have done differently?
 - [ ] Unexpected takeaways / lessons?
 - [ ] Craig: what has been the most surprising thing you have learned about k8s....during the book or otherwise?
 - [ ] Nova: Would you do it again