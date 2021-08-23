from logging import getLogger

from fastapi import APIRouter

logger = getLogger("uvicorn")
router = APIRouter()

@router.get("")
def health():
  return {"health": "OK"}