import os
from fastapi import FastAPI

app = FastAPI(title="Aegis Security Engine", version="0.1.0")

@app.get("/health")
def health_check():
    return {"status": "ok",
            "service": "aegis-sec-engine",
            "engine_version": "0.1.0"
            }

if __name__ == "__main__":
    import uvicorn
    port = int(os.getenv("PORT", 8001))
    uvicorn.run(app, host="0.0.0.0", port=port)