import xlrd

wb = xlrd.open_workbook("../data/words/word.xls")
ws = wb.sheet_by_index(0)

rows = ws.nrows
cols = ws.ncols

print(rows)
print(cols)

file = '../data/words/word'

f = open(file,'a')

for i in range(rows):
    for j in range (1,cols):
        f.write(ws.cell_value(i,j)+"\n")

f.close()