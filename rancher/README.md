# Some steps...

## 1: Preparation before presentation

### 1.1: Launch all the demo instances

```bash
# From the project's root folder, this will run the create-hosts script:
rancher/amazon/create-hosts
```

### 1.2: Configure your SSH client to access all the hosts

See `client-ssh-config.template` and edit your `~/.ssh/config` file accordingly.

### 1.3: Pre-configure the admin host

The next command will set the hostname, pull the rancher app and restart the
admin host via SSH:

```bash
ssh aws-rancher-admin '\
  sudo ros config set hostname rancher-demo ; \
  docker pull rancher/server:v0.59.0 ; \
  sudo shutdown -r now \
'
```

### 1.4: Download the agent app on each cluster node

The next command will set the hostname, pull the rancher agent and restart the
node via SSH:

```bash
# This will execute for each host:
for hostname in web-lb-01 app-01 app-02 app-03 db-01
do
  ssh aws-rancher-$hostname '\
    sudo ros config set hostname $hostname ; \
    docker pull rancher/agent:v0.9.2 ; \
    sudo shutdown -r now \
  '
done
```

## Setting hostnames on each Rancher Docker Host

Here are some examples:

```bash
# For Web Load Balancer Instances:
sudo ros config set hostname web-lb-01

# For App Host Instances:
sudo ros config set hostname app-01

# For Data Host Instances:
sudo ros config set hostname data-01
```
