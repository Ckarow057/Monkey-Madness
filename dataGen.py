import sqlite3

# Connect to the SQLite database
conn = sqlite3.connect('monkeys.db')
c = conn.cursor()

# Create the monkeys table
c.execute('''CREATE TABLE monkeys
             (monkeyID INTEGER PRIMARY KEY, 
             monkeyName TEXT, 
             monkeyBreed TEXT,
             monkeyImg TEXT, 
             monkeyFact TEXT)''')

c.execute('''CREATE TABLE userinformation
             (userID INTEGER PRIMARY KEY, 
             userSSN INTEGER, 
             userCardNumber INTEGER)''')

# Sample data to populate the table
monkeys_data = [
    (1, 'Chimpanzee', 'Pan troglodytes','TempIMG.png', 'Chimpanzees share about 98% of their DNA with humans.'),
    (2, 'Gorilla', 'Gorilla beringei','TempIMG.png' ,'Gorillas are the largest primates.'),
    (3, 'Orangutan', 'Pongo abelii','TempIMG.png', 'Orangutans are known for their distinctive red fur.'),
    (4, 'Bonobo', 'Pan paniscus','TempIMG.png', 'Bonobos are often referred to as pygmy chimpanzees.'),
    (5, 'Mandrill', 'Mandrillus sphinx','TempIMG.png', 'Mandrills are the largest monkeys in the world.'),
    # Add more data as needed
]

users_info = [
    (1, 391647654, 377456767226730),
    (2, 390372674, 377456036706215),
    (3, 390372234, 377456036723415),
    (4, 390324523, 322356036723415),
    (5, 390321233, 322356036486415)
]

# Insert the data into the table
c.executemany('INSERT INTO monkeys VALUES (?, ?, ?, ?, ?)', monkeys_data)
c.executemany('INSERT INTO userinformation VALUES (?, ?, ?)', users_info)


# Commit changes and close the connection
conn.commit()
conn.close()
