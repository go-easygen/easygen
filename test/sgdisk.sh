 # format /dev/sdb as GPT, GUID Partition Table
 sgdisk -Z /dev/sdb

 sgdisk -n 0:0:+200M -t 0:ef02 -c 0:"bios_boot" /dev/sdb
 sgdisk -n 0:0:+20G -t 0:8300 -c 0:"linux_boot" /dev/sdb
 sgdisk -n 0:0:+30G -t 0:0700 -c 0:"windows" /dev/sdb
 sgdisk -n 0:0:+10G -t 0:8200 -c 0:"linux_swap" /dev/sdb
 sgdisk -n 0:0:+12G -t 0:8300 -c 0:"os1" /dev/sdb
 sgdisk -n 0:0:+12G -t 0:8300 -c 0:"os2" /dev/sdb
 sgdisk -n 0:0:+12G -t 0:8300 -c 0:"os3" /dev/sdb
 sgdisk -n 0:0:0 -t 0:8300 -c 0:"data" /dev/sdb

 sgdisk -p /dev/sdb

 # inform the OS of partition table changes
 partprobe /dev/sdb
