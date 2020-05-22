# Episode 119 : Gatekeeper and OPA

- Hosted by @joshrosso
- 05/22/2020

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=ZJgaGJm9NJE
" target="_blank"><img src="http://img.youtube.com/vi/ZJgaGJm9NJE/hqdefault.jpg" width="480" height="360" border="10" /></a>

## Table of Contents

- [00:00:00 - Welcome to TGIK!](https://youtu.be/ZJgaGJm9NJE)
- [00:02:41 - Week in Review](https://youtu.be/ZJgaGJm9NJE?t=161)
- [00:09:50 - Admission Controller Overview](https://youtu.be/ZJgaGJm9NJE?t=590)
- [00:22:41 - OPA and Rego Overview](https://youtu.be/ZJgaGJm9NJE?t=1361)
- [00:33:81 - Gatekeeper](https://youtu.be/ZJgaGJm9NJE?t=2311)
- [01:43:16 - Wrap Up](https://youtu.be/ZJgaGJm9NJE?t=6196)
- 

## Week in Review

### Core
- Kubernetes patch release [v1.18.3](https://groups.google.com/forum/#!topic/kubernetes-announce/bBwTrUxuHDg),[1.17.6](https://groups.google.com/forum/#!topic/kubernetes-announce/EFxDXgiMLR0), [1.16.10](https://groups.google.com/forum/#!topic/kubernetes-announce/NcLzEpflUrs)
- [Monthly k8s Community Meeting](https://discuss.kubernetes.io/t/kubernetes-community-meeting-notes/35/77) Highlights:
    - Current Release Development Cycle [Taylor Dolezal - Release Lead]
      - 1.19.0-beta.0 went out on Tues, May 19th 2020
      - Enhancements FREEZE as of EOD Tues, May 19th 2020
      - Please add items to the 1.19 retro as you think of them: https://bit.ly/k8s119-retro
    - Patch Release Updates https://git.k8s.io/sig-release/releases/patch-releases.md
      - Patch releases on all branches (1.18, 1.17, 1.16) this week
      - Next patch releases likely mid-June
    - Status update from the following SIGs
      -   SIG Testing Update [Slides](https://docs.google.com/presentation/d/1H-MLhKJJVsQG2eDCEv48M_WAzMc66dKaYMgfOSGQRJM/edit#slide=id.g401c104a3c_0_0)
      -   SIG UI Update Slides [Slides](https://docs.google.com/presentation/d/1W4NioOkAF2VFiu-5t80p2vlu3_OznpugiyiViFuitaM/edit#slide=id.g338ac0a8b6_0_27)
      -   SIG API Machinery [Slides](https://docs.google.com/presentation/d/1UWRaMVtTD3yVhJ3MGBpt7LRIaRHTaQZoGlDT7Bl7jLE/edit#slide=id.g401c104a3c_0_0)
      -   SIG Usability [Slides](https://docs.google.com/presentation/d/18ISAYsExgoxk8ER47rrcv-89ZJBCZeUEdXLjT8q7HL8/)

### Community

- [Harbor 2.0 Released](https://goharbor.io/blog/harbor-2.0)
- [Draw.io vscode integration](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio)
- Nuno do Carmo and Ihor Dvoretskyi go over [setting up kind and minikube](https://kubernetes.io/blog/2020/05/21/wsl-docker-kubernetes-on-the-windows-desktop/) on WSL2/Ubuntu.  
- Caylent discusses [using OPA](https://caylent.com/leveraging-kubernetes-open-policy-agent) (Related to our topic this week!)
- That Quay outage!
- [istio 1.16](https://istio.io/news/releases/1.6.x/announcing-1.6/) has been released 

## Show Notes

- [x] Admission Control Overview
- [x] Open Policy Agent (OPA)
- [x] Gatekeeper
- [x] Conftest
- OPA Components [Architecture](https://docs.google.com/document/d/1It-Mpz36ygqrElmh2hZ3DvDIqKYyKUZN6V4d7UTlEG8/edit#heading=h.rzuko1admjwd)  and [design](https://github.com/open-policy-agent/gatekeeper/tree/master/docs/design)
- [Gatekeeper library](https://github.com/open-policy-agent/gatekeeper/tree/master/library)



## Reference Links

- [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper)
- [open-policy-agent/conftest](open-policy-agent/conftest)
- [swade1987/deprek8ion](https://github.com/swade1987/deprek8ion)
- [Webinar: K8s with OPA Gatekeeper](https://www.youtube.com/watch?v=v4wJE3I8BYM)
- [TGIK 071: OPA](https://www.youtube.com/watch?v=QU9BGPf0hBw) 
