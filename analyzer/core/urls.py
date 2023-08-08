from django.urls import path

from core.views import AnalyzerView

urlpatterns = [
    path('analyze', AnalyzerView.as_view())
]
