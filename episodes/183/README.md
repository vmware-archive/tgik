# Episode 183 : Testing Clusters via Sonobuoy: E2E tests and Your Own Plugins

- Hosted by [@johnschnake](github.com/johnschnake) and [@eleanor-millman](github.com/eleanor-millman)
- 28/01/2022

<!--- Thumbnailed embed of the video, n8Xo_ghCIOSY is the video id from the youtube url --->

<a href="https://www.youtube.com/watch?v=XreTcvKLxKI" target="_blank"><img src="http://img.youtube.com/vi/XreTcvKLxKI/hqdefault.jpg" width="480" height="360" border="10" /></a>

Photo by <a href="https://unsplash.com/@iamkaiea?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Dan Akuna</a> on <a href="https://unsplash.com/s/photos/buoy?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>

## Table of Contents

- 00:00:00 - Welcome to TGIK!
- 00:02:48 - Week in Review
- 00:07:16 - Introduction to Sonobuoy
- 00:14:35 - Running the e2e tests
- 00:37:54 - Making our own plugin
- 00:47:17 - Running our plugin
- 00:57:28 - Future plugin ideas

## Week in Review

### Core K8s

 - [Deprecation policy](https://github.com/kubernetes/website/pull/31389) change. Clarification that GA API endpoints can't be removed within a major release. This isn't a practical change, but more of a clarification of expectations.


### K8s and Cloud Native Ecosystem

 - [Policy management paper](https://www.cncf.io/blog/2022/01/28/announcing-the-kubernetes-policy-management-paper/) from CNCF Kubernetes Security Special Interest Group (SIG) and Policy Working group (WG)

## Show Notes

- Introduction to Sonobuoy
- Running the e2e tests
    - Using focus/skip
        - sonobuoy e2e
    - Using mode
        - quick
        - certified-conformance
        - non-destructive-conformance (default)
        - conformance-lite
- While its running
    - Progress/Status
    - Logs
- After completion
    - Status
    - Retrieve and untar
    - Retrieve then `results` command
- Making our own plugin
    - Requirements
        - Image to run
        - Plugin definition
        - Report results to Sonobuoy
        - Optional: Progress to Sonobuoy while running
- Plugin skeleton & plugin-helper
    - Basic boilerplate files (Docker, run script, plugin definition)
        - [Demo note: Use kind load for the image]
    - Reports results
    - Reports progress
- k8s-sigs/e2e-framework library
    - Kubernetes client
    - Hooks for before/after tests
- Running our plugin
- Future plugin ideas...
    - sonobuoy [strategy doc](https://sonobuoy.io/docs/v0.56.0/strategy/)