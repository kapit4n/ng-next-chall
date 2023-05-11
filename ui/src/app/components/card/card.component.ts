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

  hello = "Hello World"

  redirectToShowInCard() {
    if (this.onRedirect) {
      this.onRedirect(this.data.id)
    }
  }
}

