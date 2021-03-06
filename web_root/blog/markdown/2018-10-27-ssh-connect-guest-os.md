---
layout: post
title: "PuTTY로 Virtual Machine OS에 SSH 연결하기"
slug: "ssh-connect-guest-os"
date: 2018-10-02
categories: linux
tags: [linux]
---
Virtual Machine에서 guest OS를 돌릴 때 여러 불편한 점이 있습니다. 대표적으로, 지나치게 버벅거리는 현상과 host OS와의 interaction이 불편한 점 등을 들 수 있겠네요. host에서 VM으로 원격 접속하여 위 문제를 해결할 수 있습니다. 이 글에서는 다음 환경에서 실행하였습니다.

*[실행환경]*

- host OS: Windows 10 (10.0.17134)
- Virtual Machine: Oracle VirtualBox (5.2.18)
- guest OS: CentOS 7.5.1804



### Guest OS의 IP 확인

VirtualBox로 guest OS를 실행시킵니다. 터미널에서 **ifconfig **를 입력하면 다음과 같이 나옵니다.

![ifconfig](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/27/ifconfig.JPG?raw=true)



enp0s3(혹은 eth0)에서 guest IP 10.0.2.15 를 확인할 수 있습니다 (참고로, lo는 loop back interface, virbr0는 virtual network interface 입니다). guest OS의 IP를 확인했으니 위 IP 주소로 SSH 연결을 하면 될 것 같지만 연결이 되지 않습니다.

그 이유는 VirtualBox가 host를 router 삼아 NAT(Network Address Translation)을 통해 private IP(10.0.2.15) 를 만들어내기 때문입니다. 따라서 host에서 private IP로 접근하기 위해서는, router 역할을 하는 자신으로 loop back 한 다음 guest의 port로 접근해야 합니다. 이 과정에서 Port Forwarding이 필요합니다.



### VirtualBox 포트 포워딩(Port Forwarding) 설정

> *[참고] 포트 포워딩이란?*: https://opentutorials.org/course/3265/20038

*VirtualBox -설정* 에서 네트워크 탭으로 들어가면 다음과 같은 화면이 나옵니다.

![port_forwarding](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/27/port_forwarding_1.JPG?raw=true)



*고급 - 포트 포워딩* 을 클릭합니다.

![port_forwarding_2](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/27/port_forwarding_2.JPG?raw=true)



SSH 연결에 대한 포트 포워딩을 설정합니다. Host IP의 127.0.0.1은 loop back을 의미하고, guest IP의 10.0.2.15는 앞서 확인한 guest OS의 private IP 입니다. Host port는 원하시는 port를 사용하시면 되고, guest port는 SSH protocol인 22를 입력하고 저장합니다.



### PuTTY로 VM에 SSH 원격접속 하기

원격 접속을 하려면 먼저 guest OS를 실행해야 합니다. 어짜피 원격접속으로 터미널을 사용할 것이라면 헤드리스 시작으로 리소스 사용을 최소화 해줍시다.

![headless_start](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/27/headless_start.JPG?raw=true)



PuTTY를 실행시키고, 앞서 설정했던 포트 포워딩 규칙대로 접속합니다.

![putty_conn_1](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/27/putty_conn_1.JPG?raw=true)



로그인하면 guest OS의 터미널이 출력됩니다.

![putty_conn_2](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/27/putty_conn_2.JPG?raw=true)





### 참고

- [**Trying to SSH to local VM Ubuntu with Putty**](https://unix.stackexchange.com/questions/145997/trying-to-ssh-to-local-vm-ubuntu-with-putty)