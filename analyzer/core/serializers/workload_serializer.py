from rest_framework import serializers

from core.data import Workload


class WorkloadSerializer(serializers.Serializer):

    interactions = serializers.IntegerField(min_value=1)
    intervalBetweenBatchesInMilliseconds = serializers.IntegerField(min_value=0)
    batchSize = serializers.IntegerField(min_value=1)

    def validate(self, data):
        if data['batchSize'] > data['interactions']:
            raise serializers.ValidationError({'batchSize': 'it can\'t be greater than interactions.'})
        return data

    def update(self, instance, validated_data):
        raise NotImplementedError("This method is not implemented for this class")

    def create(self, validated_data):
        return Workload(
            validated_data['interactions'],
            validated_data['intervalBetweenBatchesInMilliseconds'],
            validated_data['batchSize'],
        )
