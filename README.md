# LVM prometheus exporter

Exports information about logical volumes and volume groups.

## sample output

```plaintext
# HELP lvm_lv_total_size_bytes Shows LVM LV total size in Bytes
# TYPE lvm_lv_total_size_bytes gauge
lvm_lv_total_size_bytes{lv_name="root",node="mynode",vg_name="vgubuntu"} 2.53675700224e+11
lvm_lv_total_size_bytes{lv_name="swap_1",node="mynode",vg_name="vgubuntu"} 1.023410176e+09
# HELP lvm_vg_free_bytes Shows LVM VG free size in Bytes
# TYPE lvm_vg_free_bytes gauge
lvm_vg_free_bytes{node="mynode",vg_name="vgubuntu"} 3.7748736e+07
# HELP lvm_vg_total_size_bytes Shows LVM VG total size in Bytes
# TYPE lvm_vg_total_size_bytes gauge
lvm_vg_total_size_bytes{node="mynode",vg_name="vgubuntu"} 2.54736859136e+11
```
