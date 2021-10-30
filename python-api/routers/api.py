from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from db import cruds, schemas
from db.database import get_db

router = APIRouter()

@router.get("/questionnaire/{holding_num}")
def questionnaire_by_holding_num(
  holding_num: int,
  db: Session = Depends(get_db)
):
  print("questionnaire_by_holding_num is called and holding_num is " + str(holding_num))
  if(holding_num == 0):
    return cruds.select_questionnaire_all(db=db)
  else:
    return cruds.select_questionnaire_by_holding_num(db=db, holding_num=holding_num)

@router.post("/questionnaire")
def add_questionnaire(
  questionnaire: schemas.Questionnaire,
  db: Session = Depends(get_db)
):
  return cruds.add_questionnaire_report(
    db = db,
    satisfaction_level=questionnaire.satisfaction_level,
    recommendation_level=questionnaire.recommendation_level,
    topics=questionnaire.topics,
    participation_level=questionnaire.participation_level,
    presentation_level=questionnaire.presentation_level,
    free_comment=questionnaire.free_comment,
    holding_num=questionnaire.holding_num )