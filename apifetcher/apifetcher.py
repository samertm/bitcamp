import glob
import xml.etree.ElementTree as ET
import sqlite3

def parseData(file):
    tree = ET.parse(file)
    root = tree.getroot()
    facts = None
    # TODO get name also
    for e in root.findall('pod'):
        if e.get('id') == "Input":
            name = e[0][0].text # 'plaintext' tag
        elif e.get('id') == "NutritionLabelSingle:ExpandedFoodData":
            facts = e[0][0].text # 'plaintext' tag
    food = {}
    if name is not None:
        name_tokens = name.split()
        try:
            n = name_tokens.pop(0)
            food["name"] = ""
            while n != "|":
                if len(food["name"]) != 0:
                    food["name"] += " " + n
                else:
                    food["name"] = n
                n = name_tokens.pop(0)
        except IndexError:
            # just in case...
            pass
    tokens = facts.split()
    state = "beginning"
    try:
        while True:
            t = tokens.pop(0)            
            if state == "beginning":
                if t != "serving":
                    break # error
                t = tokens.pop(0)
                if t != "size":
                    break # error
                state = "serving"
            elif state == "serving":
                if t[0] == "(":
                    # t1 is the second part
                    t1 = tokens.pop(0)
                    serv = t[1:] + " " + t1[0:len(t1)-1]
                    food["serving size"] = serv
                    state = "element"
                else:
                    t1 = tokens.pop(0)
                    # warning: terrible hack
                    if t1 == "g":
                        food["serving size"] = t + " " + t1
                        state = "element"
            elif state == "element":
                if t == "total":
                    state = "total"
                elif t == "cholesterol":
                    # do it for the rest
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["cholesterol"] = t + " " + t1
                elif t == "sodium":
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["sodium"] = t + " " + t1
                elif t == "dietary":
                    tokens.pop(0) # 'fiber'
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["dietary fiber"] = t + " " + t1
                elif t == "sugar":
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["sugar"] = t + " " + t1
                elif t == "protein":
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["protein"] = t + " " + t1
            elif state == "total":
                if t == "calories":
                    t = tokens.pop(0)
                    food["calories"] = t
                elif t == "fat":
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["fat"] = t + " " + t1
                elif t == "carbohydrates":
                    t = tokens.pop(0)
                    if t == "|":
                        continue
                    t1 = tokens.pop(0)
                    if t1 == "|":
                        continue
                    food["carbohydrates"] = t + " " + t1
                state = "element"
    except IndexError:
        # end of structure
        pass
    return food

def makeFoods():
    foods = []
    for g in glob.iglob("*.xml"):
        foods.append(parseData(g))
    return foods

def convertToInt(val):
    if val == 'NULL':
        return 'NULL'
    tokens = val.split()
    if len(tokens) == 1:
        return int(tokens[0])
    # tokens is two parts, a number and a 
    #stuffs
    if tokens[1] == 'g':
        return 1000 * int(tokens[0])
    return int(tokens[0])

prices = {
    "almond": 0.54,
    "banana": 0.36,
    "black bean": 0.66,
    "bread": 0.31,
    "brown rice": 1.10,
    "butter": 0.03,
    "carrot": 0.05,
    "cheddar cheese": 0.24,
    "egg": 0.35,
    "milk": 0.27,
    "oats": 0.88,
    "onion": 0.99,
    "peanut": 0.25,
    "prepared potato": 0.57,
    "prepared sweet pepper": 1.99,
    "prepared sweet potato": 0.72,
    "red beans": 0.08,
    "spaghetti": 0.67,
    "tomato": 1.49,
    "white rice": 0.71,
}
    
def writeDatabase():
    foods = makeFoods()
    conn = sqlite3.connect('foods.db')
    c = conn.cursor()
    c.execute('''CREATE TABLE foods
    (name text, price real, fiber integer, calories integer, sugar integer,
    fat integer, sodium integer, carbohydrates integer, serving integer,
    cholesterol integer, protein integer)''')
    for f in foods:
        vals = (f['name'],
                prices[f['name']],
                convertToInt(f.get('dietary fiber', 'NULL')),
                convertToInt(f.get('calories', 'NULL')),
                convertToInt(f.get('sugar', 'NULL')),
                convertToInt(f.get('fat', 'NULL')),
                convertToInt(f.get('sodium', 'NULL')),
                convertToInt(f.get('carbohydrates', 'NULL')),
                convertToInt(f.get('serving size', 'NULL')),
                convertToInt(f.get('cholesterol', 'NULL')),
                convertToInt(f.get('protein', 'NULL')))
        c.execute("INSERT INTO foods VALUES (?,?,?,?,?,?,?,?,?,?,?)", vals)
    conn.commit()
    conn.close()

def main():
    writeDatabase()

if __name__ == "__main__":
    main()
