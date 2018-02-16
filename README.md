# Petit VDC
This is own vdc mixed OpenStack and OpenVDC.

## Folder architecture
### compute
OpenStack の compute と同じ,
OpenVDC の Executer

---

### scheduler
OpenStack と OpenVDC の schedule と同じ 

---

### api
OpenStack の nova-api と同じ
OpenVDC は scheduler に包含されている

---

### cmd
クライアントツール
OpenStack では nova には存在しない
OpenVDC の cmd と同じ

---

### model
DB(MySQL)とのモデル とその handler

---

### template
OpenVDC の templates と同じ

---

### schema
OpenVDC の schema と同じ
