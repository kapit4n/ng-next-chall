import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  {
    path: 'subjects',
    loadChildren: () => import('./subjects/subjects.module').then(mod => mod.SubjectsModule),
  },
  {
    path: 'dashboard',
    loadChildren: () => import('./dashboard/dashboard.module').then(mod => mod.DashboardModule),
  },
  {
    path: 'categories',
    loadChildren: () => import('./categories/categories.module').then(mod => mod.CategoriesModule),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
