page_title: Installation on CentOS
page_description: Instructions for installing Docker on CentOS
page_keywords: Docker, Docker documentation, requirements, linux, centos, epel, docker.io, docker-io

# CentOS

While the Docker package is provided by default as part of CentOS-7, 
it is provided by a community repository for CentOS-6. Please note that 
this changes the installation instructions slightly between versions. 
 
These instructions work for CentOS 6 and later. They will likely work for
other binary compatible EL6 distributions such as Scientific Linux, but
they haven't been tested.

Please note that due to the current Docker limitations, Docker is able to
run only on the **64 bit** architecture.

To run Docker, you will need [CentOS6](http://www.centos.org) or higher,
with a kernel version 2.6.32-431 or higher as this has specific kernel
fixes to allow Docker to run.

## Installing Docker - CentOS-7
Docker is included by default in the CentOS-Extras repository. To install
simply run the following command.

    $ sudo yum install docker

## Installing Docker - CentOS-6
Please note that this for CentOS-6, this package is part of [Extra Packages
for Enterprise Linux (EPEL)](https://fedoraproject.org/wiki/EPEL), a community effort
to create and maintain additional packages for the RHEL distribution.

Firstly, you need to ensure you have the EPEL repository enabled. Please
follow the [EPEL installation instructions](
https://fedoraproject.org/wiki/EPEL#How_can_I_use_these_extra_packages.3F).

The `docker-io` package provides Docker on EPEL.

If you already have the (unrelated) `docker` package
installed, it will conflict with `docker-io`.
There's a [bug report](
https://bugzilla.redhat.com/show_bug.cgi?id=1043676) filed for it.
To proceed with `docker-io` installation, please remove `docker` first.

Next, let's install the `docker-io` package which
will install Docker on our host.

    $ sudo yum install docker-io

## Using Docker

Once Docker is installed, you will need to start the docker daemon.

    $ sudo service docker start

If we want Docker to start at boot, we should also:

    $ sudo chkconfig docker on

Now let's verify that Docker is working. First we'll need to get the latest
`centos` image.

    $ sudo docker pull centos

Next we'll make sure that we can see the image by running:

    $ sudo docker images centos

This should generate some output similar to:

    $ sudo docker images centos
    REPOSITORY      TAG             IMAGE ID          CREATED             VIRTUAL SIZE
    centos          latest          0b443ba03958      2 hours ago         297.6 MB

Run a simple bash shell to test the image:

    $ sudo docker run -i -t centos /bin/bash

If everything is working properly, you'll get a simple bash prompt. Type
exit to continue.

## Dockerfiles
The CentOS Project provides a number of sample Dockerfiles which you may use
either as templates or to familiarize yourself with docker. These templates
are available on github at [https://github.com/CentOS/CentOS-Dockerfiles](
https://github.com/CentOS/CentOS-Dockerfiles)

**Done!** You can either continue with the [Docker User
Guide](/userguide/) or explore and build on the images yourself.

## Issues?

If you have any issues - please report them directly in the
[CentOS bug tracker](http://bugs.centos.org).
