% DOCKER(1) Docker User Manuals
% Docker Community
% JUNE 2014
# NAME
docker-pull - Pull an image or a repository from the registry

# SYNOPSIS
**docker pull**
NAME[:TAG]

# DESCRIPTION

This command pulls down an image or a repository from the registry. If
there is more than one image for a repository (e.g., fedora) then all
images for that repository name are pulled down including any tags.
It is also possible to specify a non-default registry to pull from.

# OPTIONS
There are no available options.

# EXAMPLES

# Pull a repository with multiple images

    $ sudo docker pull fedora
    Pulling repository fedora
    ad57ef8d78d7: Download complete
    105182bb5e8b: Download complete
    511136ea3c5a: Download complete
    73bd853d2ea5: Download complete

    $ sudo docker images
    REPOSITORY   TAG         IMAGE ID        CREATED      VIRTUAL SIZE
    fedora       rawhide     ad57ef8d78d7    5 days ago   359.3 MB
    fedora       20          105182bb5e8b    5 days ago   372.7 MB
    fedora       heisenbug   105182bb5e8b    5 days ago   372.7 MB
    fedora       latest      105182bb5e8b    5 days ago   372.7 MB

# Pull an image, manually specifying path to the registry and tag

    $ sudo docker pull registry.hub.docker.com/fedora:20
    Pulling repository fedora
    3f2fed40e4b0: Download complete 
    511136ea3c5a: Download complete 
    fd241224e9cf: Download complete 

    $ sudo docker images
    REPOSITORY   TAG         IMAGE ID        CREATED      VIRTUAL SIZE
    fedora       20          3f2fed40e4b0    4 days ago   372.7 MB


# HISTORY
April 2014, Originally compiled by William Henry (whenry at redhat dot com)
based on docker.com source material and internal work.
June 2014, updated by Sven Dowideit <SvenDowideit@home.org.au>
