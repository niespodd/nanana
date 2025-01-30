# 🔓 nanana - Insecure Secret Stasher

Ever wanted to **hardcode secrets into your Go project like a total maniac**? Well, now you can! `nanana` lets you:

- ✅ Encrypt your secrets and store them in your code like a reckless genius.

- ✅ Decrypt at runtime—just type the password when prompted.

- ✅ (Optional) Auto-decrypt if the password was entered once (because convenience > security). The decryption password is stored in an environment variable (subshell), in `~/.nanana` or `[cwd]/.nanana`.

⚠️ Not safe for serious production use, but seriously safe when building PoC and not wanting to hardcode raw secrets. Anyone looking at your source code will probably laugh at this security practice (or cry).

## 📦 Installation

```bash
go get github.com/niespodd/nanana
```

## 🔥 Usage

wip