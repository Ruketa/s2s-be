import requests
import pandas as pd

def post_request_questionnaire(url: str, datafile: str):

  # 保存対象データ読込
  q_df = pd.read_csv(datafile)

  for idx, q_item in q_df.iterrows():
    # 送信用データ
    req ={
      "id": "id",
      "satisfaction_level": q_item[1],
      "recommendation_level": q_item[2],
      "topics": q_item[3],
      "participation_level": q_item[4],
      "presentation_level": q_item[5],
      "free_comment": q_item[6],
      "holding_num": 2
    }

    # postリクエスト
    with requests.post(url, json=req) as response:
      res = response.json()
      print(res)


url = "http://localhost:8000/api/questionnaire"
datafile = "questionnaire.csv"
post_request_questionnaire(url, datafile)