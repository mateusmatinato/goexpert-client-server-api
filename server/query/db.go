package query

const InitializeTable string = `
  CREATE TABLE IF NOT EXISTS exchange_rate (
  id TEXT NOT NULL,
  bid TEXT,
  created_at DATETIME NOT NULL
  );`
