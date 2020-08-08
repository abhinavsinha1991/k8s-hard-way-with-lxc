# k8s-hard-way-with-lxc


1. Provision a Ubuntu 20.04 LTS VM of 8 CPUs and 30 GB RAM.It comes pre-installed with LXD and LXC binaries(version 4.0.2 as of 8 Aug, 2020).
2. Verify if LXC version is 2.x. If yes, follow below steps, else follow video link https://www.youtube.com/watch?v=XQvQUE7tAsk directly.then jump to step 5.
3. Download this repository on your directory of choice on the VM.
4. Follow video links from step 2 in sequence to setup LXD profiles but use the  files from `kubernetes` folder of the repo you cloned.It has the recently tested ones.
5. For next steps follow video https://www.youtube.com/watch?v=NvQY5tuxALY(to understand logical flow of hard way K8s setup, skip if you know them already) and then https://www.youtube.com/watch?v=2bVK-e-GuYI.
6. Use `k8s-hard-way-steps-with-lxc.txt` file to use pre-built steps with LXC commands instead(already tested to be working).

Note: Read comments in the `k8s-hard-way-steps-with-lxc.txt` steps file wherever indicated and make appropriate changes.

Most common LXC commands are also captured in the LXC command cheatsheet file.




