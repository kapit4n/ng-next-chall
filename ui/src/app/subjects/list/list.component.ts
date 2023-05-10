import { Component, OnInit } from '@angular/core';
import { SubjectsService } from '../subjects.service'
import CardData from '../../data/card.interface'
import Transformers from '../../utils/transformers'
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

  data: CardData[] = [];
  dataSubscription: Subscription;

  constructor(private subjectsSvc: SubjectsService, private router: Router) { }

  ngOnInit() {
    this.subjectsSvc.getSubjects().subscribe(subjects => {
      this.data = subjects.map(s => Transformers.transformSubjectToCardData(s));
    });
  }

  onRedirect(subject: CardData) {
    this.router.navigate([`subjects/${subject.id}`])
  }

  ngOnDestroy() {
    this.dataSubscription ? this.dataSubscription.unsubscribe() : true;
  }
}
