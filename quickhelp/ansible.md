## 执行远程shell命令
ansible h5_api -m shell -a 'ps aux | grep qbus'

## 添加远程机器的信任
ansible-playbook push.ssh.yaml -k

## 以SUDO权限运行
ansible h5_api -m shell -a 'ps aux | grep qbus' -k -b -K

## COPY
ansible wanh5 -m copy -a 'src=/home/liuyang1-iri/tmp/wanh5 dest=/home/q/cmstpl/wanh5 owner=sync360 group=sync360' -k -b -K
