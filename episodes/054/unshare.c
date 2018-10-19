/* unshare.c

   A simple implementation of the unshare(1) command: unshare
   namespaces and execute a command.

   See https://lwn.net/Articles/531381/
*/
#define _GNU_SOURCE
#include <sched.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <sys/wait.h>

#ifndef CLONE_NEWCGROUP         /* Added in Linux 4.6 */
#define CLONE_NEWCGROUP         0x02000000
#endif

/* A simple error-handling function: print an error message based
   on the value in 'errno' and terminate the calling process */

#define errExit(msg)    do { perror(msg); exit(EXIT_FAILURE); \
  } while (0)
static void
usage(char *pname)
{
  fprintf(stderr, "Usage: %s [options] cmd [arg...]\n", pname);
  fprintf(stderr, "Options can be:\n");
  fprintf(stderr, "    -f   fork() before executing cmd "
	  "(useful when unsharing PID namespace)\n");
  fprintf(stderr, "    -C   unshare cgroup namespace\n");
  fprintf(stderr, "    -i   unshare IPC namespace\n");
  fprintf(stderr, "    -m   unshare mount namespace\n");
  fprintf(stderr, "    -n   unshare network namespace\n");
  fprintf(stderr, "    -p   unshare PID namespace\n");
  fprintf(stderr, "    -u   unshare UTS namespace\n");
  fprintf(stderr, "    -U   unshare user namespace\n");
  exit(EXIT_FAILURE);
}
int
main(int argc, char *argv[])
{
  int flags, do_fork, opt;

  flags = 0;
  do_fork = 0;
  while ((opt = getopt(argc, argv, "CfimnpuU")) != -1) {
    switch (opt) {
    case 'f': do_fork = 1;                  break;
    case 'C': flags |= CLONE_NEWCGROUP;     break;
    case 'i': flags |= CLONE_NEWIPC;        break;
    case 'm': flags |= CLONE_NEWNS;         break;
    case 'n': flags |= CLONE_NEWNET;        break;
    case 'p': flags |= CLONE_NEWPID;        break;
    case 'u': flags |= CLONE_NEWUTS;        break;
    case 'U': flags |= CLONE_NEWUSER;       break;
    default:  usage(argv[0]);
    }
  }

  if (optind >= argc)
    usage(argv[0]);

  if (unshare(flags) == -1)
    errExit("unshare");

  /* If we are unsharing the PID namespace, then the caller is *not*
       moved into the new namespace. Instead, only the children are moved
       into the namespace. Therefore, we support an option that causes
       the program to call fork() before executing the specified program,
       in order to create a new child that will be created in a new PID
       namespace. */

  if (do_fork) {
    if (fork()) {
      wait(NULL);         /* Parent waits for child to complete */
      exit(EXIT_SUCCESS);
    }

    /* Child falls through to execute command */
  }

  execvp(argv[optind], &argv[optind]);
  errExit("execvp");
}
