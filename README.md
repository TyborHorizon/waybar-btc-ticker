# Waybar BTC Tracker (Go)

Small Go script that fetches the current BTC price from CoinMarketCap.  
Main use case: show it in Waybar (or wherever you want).  

No external dependencies except your CoinMarketCap API key.

---

## Install

```bash
git clone https://github.com/your-user/waybar-btc-go.git
cd waybar-btc-go
go build -o getprice getprice.go
```

Move it somewhere in your path like:

```bash
mkdir -p ~/.local/bin
mv getprice ~/.local/bin/
```

---

## API Key Setup

You will need a free API key from [CoinMarketCap](https://coinmarketcap.com/api/).  
Either export it manually:

```bash
export COINMARKETCAP_API_KEY=your_api_key
```

Or use `.env` / `.envrc` with [direnv](https://direnv.net/)

### `.env`
```env
COINMARKETCAP_API_KEY=your_api_key
```

### `.envrc`
```bash
dotenv 
```

then allow it once:

```bash
direnv allow
```

---

## Waybar Config

waybar config block (`~/.config/waybar/config` or `.jsonc`):

```json
"custom/btc": {
  "format": "₿ ${}",
  "interval": 30,
  "exec": "~/.local/bin/getprice",
  "tooltip": false
}
```

optional css:

```css
#custom-btc {
  color: #f7931a;
  padding: 0 6px;
}
```

---

## .gitignore

dont leak your env

```gitignore
.env
.envrc
```

---

## Notes

- Script only shows BTC/USD
- Using Go builtin `net/http` and `encoding/json`
- Fast and no runtime deps
- Extend it if you want ETH, percent change, or whatever



