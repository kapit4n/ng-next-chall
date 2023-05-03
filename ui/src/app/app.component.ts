import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ui';

  constructor(private router: Router) {}

  redirectTo(to) {
    this.router.navigate([to])
  }

  onRedirectToSubject() {
    this.redirectTo("/subjects")
  }
  
  onRedirectToDashboard() {
    this.redirectTo("/dashboard")
  }
  onRedirectToCategories() {
    this.redirectTo("/categories")
  }

}
