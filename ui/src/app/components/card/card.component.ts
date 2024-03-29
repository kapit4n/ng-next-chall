import { Component, Input } from '@angular/core';
import CardData from '../../data/card.interface'

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css']
})
export class CardComponent {
  @Input('data') data: CardData;

  @Input()
  onRedirect: (id: number) => void;

  redirectToShowInCard(id: number) {
    if (this.onRedirect) {
      this.onRedirect(id)
    }
  }
}

