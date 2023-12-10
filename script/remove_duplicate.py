import csv
import argparse

parser = argparse.ArgumentParser(
    description='Remove duplicate rows from a CSV file')

parser.add_argument('-f', '--file', type=str, required=True,
                    help='CSV file to remove duplicate rows from')
parser.add_argument('-o', '--output', type=str, required=True,
                    help='CSV file to write unique rows to')

args = parser.parse_args()

unique_proxy = set()

with open(args.file, newline='') as csvfile:
    reader = csv.reader(csvfile, delimiter=':', header=None)
    for row in reader:
        unique_proxy.add(tuple(row))


with open(args.output, 'w', newline='') as csvfile:
    writer = csv.writer(csvfile, delimiter=':', header=None)
    writer.writerows(unique_proxy)
