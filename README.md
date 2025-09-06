# ctfmini (Go single-binary)

CLI เครื่องมือเล็ก ๆ สำหรับ CTF: base64, hex, url, rot, xor, vigenere, hash, hexdump, strings

## ดาวน์โหลด (พร้อมรัน)
แตก zip แล้วรันไบนารี `ctfmini` (หรือ `ctfmini.exe` บน Windows)

## ตัวอย่าง
```bash
ctfmini enc b64 "hello"             # aGVsbG8=
ctfmini enc b64 aGVsbG8= -d         # hello
ctfmini classic rot "uryyb" -n 13   # hello
ctfmini classic vigenere "attack at dawn" -k lemon   # lxfopv ef rnhr
ctfmini classic xor --in in.bin -k deadbeef --hexkey --out out.bin
ctfmini hash sum --algo sha256 --in file.zip
ctfmini util hexdump --in out.bin
ctfmini util strings --in firmware.bin --min 6
```

## สร้างไบนารีและแพ็ก zip เอง
```bash
go mod tidy
bash scripts/package.sh                # Linux/macOS
# หรือ
powershell -ExecutionPolicy Bypass -File scripts/package.ps1   # Windows
```
