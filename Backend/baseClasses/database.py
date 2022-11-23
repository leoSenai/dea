import mysql.connector
from dotenv import dotenv_values


class DB:
    
    def __init__(self, host, user, pwd, db):
        config = dotenv_values('./../.env')
        self.conn = mysql.connector.connect(
            host=config['DB_HOST'],
            user=config['DB_USER'],
            password=config['DB_PASSWORD'],
            database=config['DB_DATABASE']
            )
        self.cursor = self.conn.cursor()
    
    def doQuerySelect(self,query, values=(), addons = ()):
        if(values == () and addons == ()):
            self.cursor.execute(query)
            result = self.cursor.fetchall()
            return result
        elif(values != () and addons == ()):
            self.cursor.execute(query,values)
            result = self.cursor.fetchall()
            return result
        elif(values == () and addons != ()):
            self.cursor.execute(query+addons)
            result = self.cursor.fetchall()
            return result
        else:
            self.cursor.execute(query+addons, values)
            result = self.cursor.fetchall()
            return result