+++
date = "2018-11-30T14:59:31+11:00"
title = "Dual Boot Windows 10 & Ubuntu 18"
+++


### Dual boot Ubuntu 18 with Windows 10

> I have a couple Linux workstations, but all of them are old PC or laptop. Today I get a chance to test Ubuntu on a brand new laptop. Here I are going to write down all I did to create this dual boot laptop 

#### Caution

The laptop I worked on is Lenovo IdeaPad S model with UEFI firmware, but it doesn't mean all Lenovo laptops will work in the same way, not to mention other brand's laptop. So before you try anything, please backup all your data first. 

#### Prepare bootable USB with Ubuntu ISO

* Prepare a clean USB pen with minimum 4G volume
* Download the 18.04 LTS Desktop ISO from Ubuntu website. 
* Create a bootable USB with Rufus or any other USB build tool


#### Prepare Windows Machine for Dual-Boot

* Start the command prompt as Admin

    * Start Menu -> Command Prompt (Admin) in order to enter Windows Command Line.

* Change Window boot manager to restart the system in Safe Mode

    ```bat
    bcdedit /enum 
    
    bcdedit /set {default} safeboot minimal

    ```

* The first thing is to create a free space on the computer hard disk in case the system is installed on a single partition.


* Once in CLI, type diskmgmt.msc on prompt and the Disk Management utility should open. From here, right click on C: partition and select Shrink Volume in order to resize the partition.

    ```cmd
    C:\Windows\system32\>diskmgmt.msc
    ```

* On Shrink C: enter a value on space to shrink in MB (use at least 20000 MB depending on the C: partition size) and hit Shrink to start partition resize as illustrated below (the value of space shrink from below image is lower and only used for demonstration purposes).

Once the space has been resized you will see a new unallocated space on the hard drive. Leave it as default and reboot the computer in order to proceed with Ubuntu installation.

* Reboot the Windows to UEFI configuration

* Change the UEFI setting

    * Disable the **OS Optimized Defaults**
    * Disable the **Intel Platform Trust Technology**  (Optional)
    * Change the boot drive order 

* Plugin the USB pen and continue to boot the system

#### Install Ubuntu

* Install Ubuntu Desktop as usual
* Recommand to install minimal version. You always can install other applications later.


#### Rollback the Windows Safe Mode

* Launch command promt as Admin and run following command

```
bcdedit /deletevalue {default} safeboot
```

#### Make your life easier

- Ubuntu

    * Update fstab to add the Windows mount point at system start up
    * Update the Grub file if you want to boot Windows as default OS

- Windows

    * Install the Linx Reader to access the Ubuntu files from Windows
    
     