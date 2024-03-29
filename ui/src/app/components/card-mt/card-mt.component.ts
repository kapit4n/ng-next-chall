import { Component, Input } from '@angular/core';
import CardData from '../../data/card.interface'

@Component({
  selector: 'app-card-mt',
  templateUrl: './card-mt.component.html',
  styleUrls: ['./card-mt.component.css']
})
export class CardMtComponent {
  @Input('data') data: CardData;
}
