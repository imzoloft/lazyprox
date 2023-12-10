import pandas as pd
import argparse

parser = argparse.ArgumentParser(
    description='Remove duplicate rows from one CSV file to antoher')

parser.add_argument('-f', '--file', type=str, required=True,
                    help='CSV file that contains the lines to remove from the other file')
parser.add_argument('-f1', '--file1', type=str, required=True,
                    help='CSV file to remove lines in')

args = parser.parse_args()

df1 = pd.read_csv(args.file, delimiter=':', header=None)
df2 = pd.read_csv(args.file1, delimiter=':', header=None)

df2 = df2[~df2.isin(df1)].dropna()

df2.to_csv('proxies_modified.csv', index=False, sep=':', header=None)
