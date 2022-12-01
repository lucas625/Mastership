from django.urls import path

from core.views import AnalyserView

urlpatterns = [
    path('png', AnalyserView.as_view())
]
