from django.urls import path

from core.views import PingView
from core.views import PingCalculatorView

urlpatterns = [
    path('ping', PingView.as_view()),
    path('ping-calculator', PingCalculatorView.as_view())
]
