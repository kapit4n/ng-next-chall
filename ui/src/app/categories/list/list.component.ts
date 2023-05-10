import { Component, OnInit } from '@angular/core';
import { CategoriesService } from '../categories.service';
import CardData from '../../data/card.interface'
import Transformers from '../../utils/transformers'
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

  data: CardData[] = [];
  dataSubscription: Subscription;
  
  constructor(private catService: CategoriesService) {

  }

  ngOnInit() {
    this.catService.getAllCategories().subscribe(cats => {
      this.data = cats.map(c => Transformers.transformCategoryToCardData(c));
    })
  }

  ngOnDestroy() {
    this.dataSubscription ? this.dataSubscription.unsubscribe() : true;
  }
}
