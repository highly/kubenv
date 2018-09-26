# kubenv
Easy way to change kubernetes environment by switching config file 

### Installation

```text
unzip kubenv-xxx.zip

chmod +x kubenv

mv kubenv /usr/local/bin
```

### Usage

```text
USAGE:  kubenv -e <extension name of config file>
```

### Example


```text
mkdir $HOME/.kube

mv config.prod config.uat config.beta-qa $HOME/.kube

// to production environment
kubenv -e prod

// to qa environment

kubenv -e beta-qa
```
