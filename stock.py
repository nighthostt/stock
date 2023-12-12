from jqdatasdk import *
import jqdatasdk as jq
import pandas as pd
jq.auth('15356157529', 'Tww12345')
#df = jq.get_price("000001.XSHE")
#print(df)
data = jq.get_all_securities(types=['stock'])

df = pd.DataFrame(data)
df.to_csv("stock.csv", encoding='utf-8', index=True, header=False)
