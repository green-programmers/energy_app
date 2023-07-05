import psycopg2

class BasicCommands:
    def __init__(self,dbname , username , password ,cursor = None):
        self.dbname = dbname
        self.username = username
        self.password = password
        self.cursor = cursor

    def connect(self):
        self.cursor = psycopg2.connect(f"dbname={self.dbname} user={self.username} password={self.password}").cursor()

    def select_query(self,table):
        self.cursor.execute(f"SELECT * FROM {table}")

    def fetch_all(self , num_of_rows):
        fetch = self.cursor.fetchall()
        inc = 0
        for row in fetch:
            while inc < num_of_rows:
                print(row[inc])
                inc+=1
