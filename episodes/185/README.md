# Episode 185 --  Cartographer: the K8s-native supply chain choreographer

- Hosted by [David Espejo](https://twitter.com/davidmirror)
    - Guests: 
        - [Rasheed Abdul-Aziz](https://github.com/squeedee) 
        - [Waciuma Wanjohi](https://github.com/waciumawanjohi)
-  11/02/2022

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=TJPGn0-hpPY" target="_blank"><img src="https://i9.ytimg.com/vi/TJPGn0-hpPY/mqdefault.jpg?v=62041e08&sqp=CLiQqpAG&rs=AOn4CLCnFZgZfrbcbJLS5qdZkk1Zf4UK-w" width="380" height="220"  border="10" /></a>

Photo by <a href="https://unsplash.com/@teckhonc?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">T.H. Chia</a> on <a href="https://unsplash.com/?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
  

## Table of Contents

- 00:00:36 - Welcome to TGIK
- 00:01:35 - Introductions
- 00:04:58 - Week in review
- 00:12:42 - Why Cartographer?
- 00:32:37 - Demo: basic supply chain
- 01:12:28 - Wrapping up

## Week in Review

### Core K8s

 - Enhancements freeze for K8s 1.24 in effect since Feb 4th. 
     - **25** enhancements removed from milestone: https://docs.google.com/spreadsheets/d/1T21mUTvPh70NB2eseHjCyD4LgRjyxWI9Bd1SoP8zAwA/edit?usp=sharing
- New blog post: spotlight on SIG Multicluster https://kubernetes.io/blog/2022/02/07/sig-multicluster-spotlight-2022/
- K8s community Office Hours (SIG Contribex edition) will be on Wed Feb 16th: https://github.com/kubernetes/community/blob/master/events/office-hours.md 


### K8s and Cloud Native Ecosystem
- Testify Sec turned the Witness repo into public: https://github.com/testifysec/witness
- Cluster API v1.1.0 released https://github.com/kubernetes-sigs/cluster-api/releases/tag/v1.1.0
- Continuous Delivery Foundation Conference final CFP deadline is Feb 18th: https://cd.foundation/blog/2022/01/10/cdcon-2022-call-for-papers-open/
- LFX mentoring program deadline is Feb 18th: https://www.cncf.io/blog/2022/02/09/lfx-spring-2022-internships-are-open-apply-for-cncf-projects-by-february-13th-%ef%bf%bc/
- CNCF TOC has approved the archiving of the OpenTracing project: https://github.com/cncf/toc/pull/710?utm_source=hs_email&utm_medium=email&_hsenc=p2ANqtz-9kvZxI8kHdtnPXesigFg-hxk6s4k0gvPFo49q4LFdBAtqmlwaF4YYiAr5NN_03ieIQiEU0NZgxSQSVD0N3fYVI_tAzuw


## Show notes

- Introduction to [Cartographer](https://cartographer.sh/) @Rasheed [#574](https://github.com/vmware-tanzu/cartographer/pull/574)
    - [12 factor apps](https://12factor.net) Heroku push changed the world!
- Demo @Waciuma
    - Heroku push experience helped reduce cognitive load for devs in the past, then we moved to K8s... https://static.sched.com/hosted_files/cloudnativedevxdayna21/39/DevX%20Day%20-%20KubeCon%20NA%202021%20Slides_%20From%20Kubernetes%20to%20PaaS%20to%20Developer%20Control%20Planes.pdf 
    - Architecture diagrams: https://cartographer.sh/docs/v0.2.0/architecture/
    - Cartographer examples folder: 