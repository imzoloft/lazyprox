## <h4>Utilitaries</h4>

---

Remove duplicated lines in the same file. <br/>
Remove lines that are present in one file in another one.

---

<h5>Options</h5>
<b>Remove Duplicate</b> <br/>
-f, --file: An input file <br/>
-o, --ouput: An output file <br/>

python3 remove_duplicate.py -f input.csv -o output.csv <br/><br/>

<b>Remove valid lines</b> <br/>
-f, --file: An input file with the lines you want to delete in the other file <br/>
-o, --ouput: An input file where the lines will be deleted <br/>

python3 remove_valid_lines.py -f input.csv -f1 input1.csv
