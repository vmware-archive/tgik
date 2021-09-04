# Episode 164: The K8s Pod Specification ~ RuntimeClasses, Hooks, Cgroups, and whatever else we run into.

```
Jay Vyas, Duffie Coolie, and Nishad Mathur as we look at
the evolving data model for pods: container lifecycle events, hooks, init pods, hostProcess containers, and other heebie-geebies in the "what's a container" domain.
```
@jayunit100 
@mauilion

## Week in Review
- https://github.com/kubernetes/kubernetes/issues/104648 Interesting kubelet regression
- https://hackmd.io/5rrv64_QQzyhnzf15mZ73Q blog post from @perithompson on hostProcess containers
- https://www.manning.com/books/kubernetes-on-windows  

- https://ebpf.io/summit-2021/ checkout mauilion and lizrice hosting ebpf.io summit 2021
  - cillium
    -  replace kube proxy for routing
    -  5.4+ 
## Show Notes!

- https://github.com/kubernetes/kubernetes/issues/104648
- https://www.youtube.com/watch?v=u8h0e84HxcE
- Do we know why the systemd driver went away?
- SigHup... SigKill
- Whats the diff between kill -1 vs -9 ?
    - `man signal`
    - https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
    - https://stackoverflow.com/questions/43724467/what-is-the-difference-between-kill-and-kill-9#:~:text=Killing%20a%20process%20using%20kill,and%20kills%20the%20process%20immediately.

## PIDs
- disconnecting the parent
    - `get.docker.io` <-- from now on @nishad 
    - 
### How to create a PostStart hook
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: lifecycle-demo
spec:
  containers:
  - name: lifecycle-demo-container
    image: nginx
    lifecycle:
      postStart:
        exec:
          command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
      preStop:
        exec:
          command: ["/bin/sh","-c","nginx -s quit; while killall -0 nginx; do sleep 1; done"]
```
- saw a exit 137  failed to pre-stop hook:
    - 128+9 = SIGKILL -thanks @Nicolas D 
    - Overran grace period

### How to create a PreStart hook
https://github.com/kubernetes/enhancements/issues/753
https://github.com/kubernetes/enhancements/pull/2869/

- `kubectl f l u`

# Cgroups
- `ls -al /proc/PID/ns/`: 
- `systemd-cgls`:
- cgroup controller
    - important docker and kublet agree on cgroups allocator
    - systemd or cgroupfs
    - both handle slicing resources between cgroups 
        - so if you dont pick the right one both will assign the same 1g of memory 2 different cgroups
        - may hit an OOM  where reserved memory isnt really reserved



- 0:00 week in review
- 1:00 surprise guest , tgik Vet mauilion from Isovalent joins us!
- 5:00 HostProcess container blog post by perithompson
- 5:31 Kubernets on Windows book
- 7:00 Ebpf summit, netns cookies, routing improvments
- 11:00 Nishad introduces his nested docker+celery problem
- 12:20 "Its friday were gonna do things right, said nobody ever"
- 15:30 PIDs inside a kind cluster is weird
- 19:15 Kernel tasks have weird names
- 23:00 containerd shim synthetic parent to the pause/runsvcdir startups,
- 24:00 pause container holding namespaces for sharing
- 25:00 containers sharing PID namespaces
- 32:00 runc creating the process for containerd/dockershim
- 33:00 containerd-shim as the holder for runc's process
- 48:00 After TerminationGracePeriods, SIGKILL 
- 49:30 SIGTERM sent first before the Kill happens
- 56:00 Init pods, pods, processes, postStart, preStop hooks
- 58:00 no gaurantee that the preStop will finish
- 1:01:00 container lifecycle hooks docs page 
- 1:07:00 kind cluster became a zombie, killing mirror pods is meaningless
- 1:08:00 "Shadow pods" oops mirror pod
- 1:09:00 trying to fix my borked kind cluster
- 1:14:00 Fei joins us 
- 1:17:00 kindnet , dualstack
- 1:20:00 systemctl -flu 
- 1:23:00 cillium port ranges pleeease so we can graduate networkpolicy API improvements to GA !
- 1:32:00 containers in the same pod have DIFFERENT pid namespaces normally unless you configure explicitly
- 1:35:00 CGroupFS vs Systemd allocator 
- 1:39:00 CgroupFS controller doesnt know about space allocated by others
