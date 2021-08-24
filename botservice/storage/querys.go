package storage

func (res *Respuestas) GetMsg(cmd string) (*Respuestas, error) {
	db := MySqlConn()
	q := `SELECT id, comando, respuesta FROM messages WHERE comando = ?`

	err := db.QueryRow(q, cmd).Scan(&res.ID, &res.Comando, &res.Respuesta)
	defer db.Close()
	if err != nil {
		return &Respuestas{}, err
	}

	return res, nil
}
