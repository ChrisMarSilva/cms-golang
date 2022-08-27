package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// go mod init github.com/chrismarsilva/cms.golang.tnb.coinmarketcap
// go mod tidy

// go run main.go

func main() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "<TOKEN>")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

}

/*
url = 'https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=5000'
headers = { 'Accepts': 'application/json', 'X-CMC_PRO_API_KEY': '<TOKEN>', }   # meu
session = Session()
session.headers.update(headers)
try:
    response = session.get(url)
    data = json.loads(response.text)
except (ConnectionError, Timeout, TooManyRedirects) as e:
    print(e)



# This creates a long string of all the top 100 crypto currency symbols.
symbolstr=','.join(('BTC,ETH,BNB,XRP,USDT,ADA,DOT,UNI,LTC,LINK,XLM,BCH,THETA,FIL,USDC,TRX,DOGE,WBTC,VET,SOL,KLAY,EOS,XMR,LUNA,MIOTA,BTT,CRO,BUSD,FTT,AAVE,BSV,XTZ,ATOM,NEO,AVAX,ALGO,CAKE,HT,EGLD,XEM,KSM,BTCB,DAI,HOT,CHZ,DASH,HBAR,RUNE,MKR,ZEC,ENJ,DCR,MKR,ETC,GRT,COMP,STX,NEAR,SNX,ZIL,BAT,LEO,SUSHI,MATIC,BTG,NEXO,TFUEL,ZRX,UST,CEL,MANA,YFI,UMA,WAVES,RVN,ONT,ICX,QTUM,ONE,KCS,OMG,FLOW,OKB,BNT,HNT,SC,DGB,RSR,DENT,ANKR,REV,NPXS,VGX,FTM,CHSB,REN,IOST,CELO,CFX'))
symbolstr
symbol_list=symbolstr.split(',') # Makes symbolstr into a list for later for loop
symbol_list[:5]
url = f'https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest'
headers = {'Accepts': 'application/json',    'X-CMC_PRO_API_KEY': '<TOKEN>',}# meu
parameters = { 'symbol': symbolstr,'convert':'BRL'  }
session = Session()
session.headers.update(headers)
try:
    response = session.get(url, params=parameters)
    data1 = json.loads(response.text)
except (ConnectionError, Timeout, TooManyRedirects) as e:
    data1 = json.loads(response.text)


	print('id=',data1['data']['AAVE']['id'])
print('cmc_rank=',data1['data']['AAVE']['cmc_rank'])
print('last_updated=',data1['data']['AAVE']['last_updated'])
print('name=',data1['data']['AAVE']['name'])
print('slug=',data1['data']['AAVE']['slug'])
print('symbol=',data1['data']['AAVE']['symbol'])
print('last_updated=',data1['data']['AAVE']['quote']['BRL']['last_updated'])
print('market_cap=',data1['data']['AAVE']['quote']['BRL']['market_cap'])
print(f"percent_change_30d={data1['data']['AAVE']['quote']['BRL']['percent_change_30d']:.2f}")
print(f"percent_change_7d={data1['data']['AAVE']['quote']['BRL']['percent_change_7d']:.2f}")
print(f"percent_change_24h={data1['data']['AAVE']['quote']['BRL']['percent_change_24h']:.2f}")
print(f"price={data1['data']['AAVE']['quote']['BRL']['price']:.2f}")


with open('coinmap.txt', 'w') as f:
    for symbol in symbol_list:
        cid=data1['data'][symbol]['id']
        name=data1['data'][symbol]['name']
        price=data1['data'][symbol]['quote']['BRL']['price']
        line=f'{cid}, {name}, {symbol}, {price}'
        print(line)
        f.write(f'{line}\n')


url = 'https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest'
parameters = {'start':'1','limit':'5000','convert':'BRL'}
headers = {'Accepts': 'application/json',  'X-CMC_PRO_API_KEY': '<TOKEN>',}# meu
session = Session()
session.headers.update(headers)
try:
    response = session.get(url, params=parameters)
    data2 = json.loads(response.text)
except (ConnectionError, Timeout, TooManyRedirects) as e:
    data2 = json.loads(response.text)


	data2['status']
# data_status = data['status']
# api_timestamp = data_status['timestamp'] # o carimbo de data/hora do ponto de preço
# api_credits = data_status['credit_count'] # o número de créditos gastos na última solicitação


for row in data2['data']:
    print('id=',row['id'], 'cmc_rank=',row['cmc_rank'], 'name=',row['name'], 'symbol=',row['symbol'], 'last_updated=',row['quote']['BRL']['last_updated'], f"percent_change_24h={row['quote']['BRL']['percent_change_24h']:.2f}", f"price={row['quote']['BRL']['price']:.2f}")

# data_quote = data['data'][0]['quote']['BRL']
# price = np.round(data_quote['price'], 1) # the price
# volume_24 = np.round(data_quote['volume_24h'], 1) # the volume in the last 24 hours


*/
