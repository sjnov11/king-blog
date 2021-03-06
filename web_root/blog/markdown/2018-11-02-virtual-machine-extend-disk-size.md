---
layout: post
title: "가상머신 디스크 공간 늘이기"
slug: "virtual-machine-extend-disk-size"
date: 2018-11-02
categories: Linux
tags: [linux]
---
VirtualBox는 기본적으로 guest OS의 디스크 용량을 8GB로 동적 할당합니다. 동적 할당 설정을 해 둔다면 guest OS의 사용량만큼 host의 디스크를 사용하기 때문에 VirtualBox의 설정을 가급적 크게 설정하는 것을 권합니다. 

충분한 크기의 디스크 용량으로 설정하지 않았을 경우, guest OS에서의 디스크 공간이 부족해 이를 확장해야하는 경우가 발생합니다. 이 때, 어떻게 guest OS의 디스크 공간 늘이는 지 그 방법에 대해서 알아보겠습니다.

### 실행환경

- host OS: Windows 10 (10.0.17134)
- Virtaul Machine: Oracle VirtualBox (5.2.18)
- guest OS: CentOS (7.5.1804)

### Host 작업

1. Host OS에서는 VirtualBox의 이미지 파일의 크기를 키워주면 됩니다. cmd 창에서 아래와 같이 입력을 합니다.

   ```bash
   > [VBoxManage.exe] modifyhd [guestOS_Image.vdi] --resize [size(MB)]
   ```

   예를 들어, VirtualBox를 기본경로로 설치, 40GB로 늘이려면 다음과 같습니다.

   ```bash
   > "C:\Program Files\Oracle\VirtualBox\VBoxManage.exe" modifyhd "C:\Users\사용자_이름\VirtualBox VMs\centos-test\centos-test.vdi" --resize 40960
   ```


2. VirtualBox를 재실행하면 이미지 파일의 크기가 변경된 것을 확인할 수 있습니다.

   ![host_2](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/11/02/host_2.png?raw=true)



### Guest 작업

Host에서 Image 크기를 키워준만큼, guest OS에서 파티션 영역을 키워주어야 합니다.

1. ``` fdisk -l``` 를 입력하여 디스크 크기가 늘어났는지 확인합니다. 

   ![guest_1](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/11/02/guest_1.png?raw=true)

   참고로 /dev/mapper/ 는 LVM의 logical volume을 나타냅니다. **Physical volume은 /dev/sda 이고, 이것이 실제 파티션이 생성되는 디스크**입니다.

2. ```df -h``` 를 입력하여 **파일 시스템에 mapping 되는 logical volume**을 확인합니다. (/dev/mapper/centos-root)

   ![guest_2](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/11/02/guest_2.PNG?raw=true)

3. ```fdisk /dev/sda``` 을 입력하여 physical volume을 파티션 설정을 변경합니다.

   1) ```p``` 를 입력 하여 현재 partition table을 출력합니다.

   2) ```d``` 와 ```2``` 를 순서대로 입력하여 boot partition이 아닌 partition을 table에서 삭제합니다. **이는 partition을 새로 설정하기 위해 지우는 것이지 partition 자체를 지우는 것이 아닙니다.** 또, **boot partition을 지우지 않도록 주의**합니다.

   3) ```n``` , ```p ``` ,```2``` 를 순서대로 입력하여 지웠던 2번의 partition을 다시 생성합니다.

   4) First sector 값을 설정합니다. Disk에서 비워져있는 값의 처음을 default로 합니다. ```enter``` 를 누르면 default로 선택됩니다.

   5) Last sector 값을 설정합니다. Disk 크기의 맨 마지막 값을 default로 합니다. ```enter``` 를 누르면 default로 선택됩니다.

   6) ```t``` 를 입력합니다. ```2``` 를 입력하여 2번 partition을 선택, ```8e``` 를 입력하여 partition type을 Linux LVM으로 설정합니다.

   7) ```p``` 를 입력해보면 1에서 출력한 partition의 start-end 보다 더 커져있는 것을 확인할 수 있습니다. 우리는 이렇게 partitioning table을 다시 설정한 것입니다.

   ![guest_3-7](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/11/02/guest_3-7.PNG?raw=true)

   6) ```w``` 를 입력하여 수정한 **partition table을 저장합니다.**

   7) ```reboot``` 하여 재부팅합니다.

4. Physical Volume의 실제 크기를 partition table에서 설정한 크기만큼 조절합니다. ```pvresize /dev/sda2``` 를 입력합니다. ```pvscan``` 을 입력하면 physical volume의 크기가 변한 것을 확인할 수 있습니다.

   ![guest_4](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/11/02/guest_4.PNG?raw=true)

5. 증가한 Physical Volume 을 가지고 Logical Volume에 mapping 합니다.

   ```lvextend -l +100%FREE /dev/mapper/centos-root``` 

6. Logical Volume의 크기를 증가한 크기만큼 늘여줍니다.

   ```lvresize -r -l +100%FREE /dev/mapper/centos-root``` 

7. ```df -h``` 를 통해 늘어난 공간을 확인합니다.



### 참고

- [**Resize a Hard Disk for a Virtual Machine**](https://gist.github.com/christopher-hopper/9755310)
- **[VirtualBox 기반에서 CentOS7 저장소 용량 늘리는 방법](http://wangin9.tistory.com/entry/vbox-centos7-memory)**

