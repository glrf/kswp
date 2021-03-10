# kswp

One way to manage the access to multiple Kubernetes clusters

## Motivation

Managing kubeconfigs for multiple Kubernetes clusters can quickly become confusing. While kubeconfig was designed to 
configure access to multiple clusters with multiple contexts, if you're anything like me you quickly loose track of what
cluster you're actually working on and will make mistakes.

This lead to me maintaining multiple kubeconfig files and moving them manually. Obviously this lead to more accidents.
This is why I developed `kswp`. With `kswp` you provide a list workspaces, each with their own separate kubeconfig, and
you can quickly switch between them. 

This is essentially the same work flow of maintaining a separate config file for each environment, but the more automated
switching lead to less errors in my case.

## Usage

First you will need to provide a configuration file. The configuration file needs to be called `kswp.yml` and needs
to either be located at `~/.kube` or `~/.config/kswp`. It contains the list of environment you need and the location
of their config file

    configs:
      - name: production
        path: /home/bob/.kube/gke.yaml
      - name: staging
        path: /home/bob/.kube/stage.yml
      - name: local
        path: /home/bob/.kube/microk8s.config

With that in place you can switch to the staging environment with

    kswp staging 

This will overwrite your `~/.kube/config`


## Installation

To install this you can simply run 

    go get -u github.com/glorfischi/kswp/cmd/kswp


## Autocomplete

`kswp` supports auto-completion. Follow the instructions for your shell to enable it

### Bash

    $ source <(kswp completion bash)

    # To load completions for each session, execute once:
    # Linux:
    $ kswp completion bash > /etc/bash_completion.d/kswp
    # macOS:
    $ kswp completion bash > /usr/local/etc/bash_completion.d/kswp

### Zsh

    # If shell completion is not already enabled in your environment,
    # you will need to enable it.  You can execute the following once:

    $ echo "autoload -U compinit; compinit" >> ~/.zshrc

    # To load completions for each session, execute once:
    $ kswp completion zsh > "${fpath[1]}/_kswp"

    # You will need to start a new shell for this setup to take effect.

### fish:

    $ kswp completion fish | source

    # To load completions for each session, execute once:
    $ kswp completion fish > ~/.config/fish/completions/kswp.fish
