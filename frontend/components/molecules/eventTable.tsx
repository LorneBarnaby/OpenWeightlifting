import {
  Table,
  TableHeader,
  TableColumn,
  TableCell,
  TableRow,
  TableBody, Link,
} from '@nextui-org/react'
import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import {AsyncListData, useAsyncList} from "@react-stately/data";
import { SortDescriptor} from "@react-types/shared";

export const EventTable = ({
                               history,
                             }: {
  history: LifterResult[]
}) => {

  let list = useAsyncList({
    async load () {
      return {items: history}
    },
    async sort({items, sortDescriptor}:{items:LifterResult[], sortDescriptor:SortDescriptor}) {
      return {
        items: items.sort((a, b) => {
          let col = (sortDescriptor.column?? "sinclair") as keyof LifterResult;  
          let first = parseInt(a[col].toString());
          let second = parseInt(b[col].toString());

          let cmp = first < second ? -1 : 1;

          // Rank misses at bigger weights higher
          if(first < 0 && second < 0){
            cmp *= -1
          }
          if (sortDescriptor.direction === "descending") {
            cmp *= -1;
          }
          return cmp;
        }),
      };
    },
  });

  return (
    <Table onSortChange={list.sort} sortDescriptor={list.sortDescriptor}>
      <TableHeader>
        <TableColumn>Date</TableColumn>
        <TableColumn>Name</TableColumn>
        <TableColumn allowsSorting key="bodyweight">Bodyweight</TableColumn>
        <TableColumn allowsSorting key="snatch_1">1st Snatch</TableColumn>
        <TableColumn allowsSorting key="snatch_2">2nd Snatch</TableColumn>
        <TableColumn allowsSorting key="snatch_3">3rd Snatch</TableColumn>
        <TableColumn allowsSorting key="cj_1">1st C&J</TableColumn>
        <TableColumn allowsSorting key="cj_2">2nd C&J</TableColumn>
        <TableColumn allowsSorting key="cj_3">3rd C&J</TableColumn>
        <TableColumn allowsSorting key="total">Total</TableColumn>
        <TableColumn allowsSorting key="sinclair">Sinclair</TableColumn>
      </TableHeader>
      <TableBody>
        {list.items.map((lift, index) => {
          const {
            date,
            lifter_name,
            bodyweight,
            snatch_1,
            snatch_2,
            snatch_3,
            cj_1,
            cj_2,
            cj_3,
            total,
            sinclair,
          } = lift

          const lifter_page = '../lifter?name=' + lifter_name

          return (
            <TableRow key={`history-${index}`}>
              <TableCell>{date}</TableCell>
              <TableCell>
                <Link href={lifter_page}>
                  {lifter_name}
                </Link>
              </TableCell>
              <TableCell>{bodyweight}</TableCell>
              <TableCell>{snatch_1}</TableCell>
              <TableCell>{snatch_2}</TableCell>
              <TableCell>{snatch_3}</TableCell>
              <TableCell>{cj_1}</TableCell>
              <TableCell>{cj_2}</TableCell>
              <TableCell>{cj_3}</TableCell>
              <TableCell>{total}</TableCell>
              <TableCell>{sinclair}</TableCell>
            </TableRow>
          )
        })}
      </TableBody>
    </Table>
  )
}
