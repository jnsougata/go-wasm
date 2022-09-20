from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.responses import FileResponse

app = FastAPI()

app.mount("/static", StaticFiles(directory="./static"), name="static")

@app.get("/")
def read_root():
    return FileResponse("templates/index.html")