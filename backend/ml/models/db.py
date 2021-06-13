import mysql.connector
import config
import time

def send_db(data):
    mydb = mysql.connector.connect(
        host=config.HOST,
        user=config.USER,
        password=config.PASS,
        database=config.DB
    )
    mycursor = mydb.cursor()
    sql = "INSERT INTO children_texts (children_id, text, timestamp) VALUES (%s, %s, %s)"
    mycursor.executemany(sql, data)
    mydb.commit()
    print(mycursor.rowcount, "was inserted")


